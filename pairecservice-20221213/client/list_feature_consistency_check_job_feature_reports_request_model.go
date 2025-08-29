// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureConsistencyCheckJobFeatureReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListFeatureConsistencyCheckJobFeatureReportsRequest
	GetInstanceId() *string
	SetLogItemId(v string) *ListFeatureConsistencyCheckJobFeatureReportsRequest
	GetLogItemId() *string
	SetLogRequestId(v string) *ListFeatureConsistencyCheckJobFeatureReportsRequest
	GetLogRequestId() *string
	SetLogUserId(v string) *ListFeatureConsistencyCheckJobFeatureReportsRequest
	GetLogUserId() *string
}

type ListFeatureConsistencyCheckJobFeatureReportsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9010
	LogItemId *string `json:"LogItemId,omitempty" xml:"LogItemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// F7AC05FF-EDE7-5C2B-B9AE-33D6DF4178BA
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1010
	LogUserId *string `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
}

func (s ListFeatureConsistencyCheckJobFeatureReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobFeatureReportsRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsRequest) GetLogItemId() *string {
	return s.LogItemId
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsRequest) GetLogRequestId() *string {
	return s.LogRequestId
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsRequest) GetLogUserId() *string {
	return s.LogUserId
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsRequest) SetInstanceId(v string) *ListFeatureConsistencyCheckJobFeatureReportsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsRequest) SetLogItemId(v string) *ListFeatureConsistencyCheckJobFeatureReportsRequest {
	s.LogItemId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsRequest) SetLogRequestId(v string) *ListFeatureConsistencyCheckJobFeatureReportsRequest {
	s.LogRequestId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsRequest) SetLogUserId(v string) *ListFeatureConsistencyCheckJobFeatureReportsRequest {
	s.LogUserId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsRequest) Validate() error {
	return dara.Validate(s)
}
