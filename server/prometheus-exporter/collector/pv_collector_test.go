/*
 *  Copyright (c) Huawei Technologies Co., Ltd. 2023-2023. All rights reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package collector

import (
	"reflect"
	"testing"
)

func Test_parsePVStorageIDGetName(t *testing.T) {
	// arrange
	mockInDataKey := ""
	mockMetricsName := ""
	mockInData := map[string]string{
		"NAME": "fake_name",
		"ID":   "fake_data",
	}

	// action
	got := parsePVStorageID(mockInDataKey, mockMetricsName, mockInData)

	// assert
	if !reflect.DeepEqual(got, "fake_data") {
		t.Errorf("parseStorageData() got = %v, want %v", got, "fake_data")
	}
}

func Test_parsePVStorageIDGetObjectName(t *testing.T) {
	// arrange
	mockInDataKey := ""
	mockMetricsName := ""
	mockInData := map[string]string{
		"ObjectName": "fake_name",
		"ObjectId":   "fake_data",
	}

	// action
	got := parsePVStorageID(mockInDataKey, mockMetricsName, mockInData)

	// assert
	if !reflect.DeepEqual(got, "fake_data") {
		t.Errorf("parseStorageData() got = %v, want %v", got, "fake_data")
	}
}

func Test_parsePVCapacityUsageSan(t *testing.T) {
	// arrange
	mockInDataKey := ""
	mockMetricsName := ""
	mockInData := map[string]string{
		"sbcStorageType": "oceanstor-san",
		"CAPACITY":       "100",
		"ALLOCCAPACITY":  "10",
	}

	// action
	got := parsePVCapacityUsage(mockInDataKey, mockMetricsName, mockInData)

	// assert
	if !reflect.DeepEqual(got, "10.00") {
		t.Errorf("parseStorageData() got = %v, want %v", got, "fake_data")
	}
}
