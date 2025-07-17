// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartEnvironmentFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *RestartEnvironmentFeatureResponseBody
	GetCode() *int32
	SetData(v string) *RestartEnvironmentFeatureResponseBody
	GetData() *string
	SetMessage(v string) *RestartEnvironmentFeatureResponseBody
	GetMessage() *string
	SetRequestId(v string) *RestartEnvironmentFeatureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestartEnvironmentFeatureResponseBody
	GetSuccess() *bool
}

type RestartEnvironmentFeatureResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2983BEF7-4A0D-47A2-94A2-8E9C5E63****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartEnvironmentFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartEnvironmentFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *RestartEnvironmentFeatureResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RestartEnvironmentFeatureResponseBody) GetData() *string {
	return s.Data
}

func (s *RestartEnvironmentFeatureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RestartEnvironmentFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartEnvironmentFeatureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestartEnvironmentFeatureResponseBody) SetCode(v int32) *RestartEnvironmentFeatureResponseBody {
	s.Code = &v
	return s
}

func (s *RestartEnvironmentFeatureResponseBody) SetData(v string) *RestartEnvironmentFeatureResponseBody {
	s.Data = &v
	return s
}

func (s *RestartEnvironmentFeatureResponseBody) SetMessage(v string) *RestartEnvironmentFeatureResponseBody {
	s.Message = &v
	return s
}

func (s *RestartEnvironmentFeatureResponseBody) SetRequestId(v string) *RestartEnvironmentFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartEnvironmentFeatureResponseBody) SetSuccess(v bool) *RestartEnvironmentFeatureResponseBody {
	s.Success = &v
	return s
}

func (s *RestartEnvironmentFeatureResponseBody) Validate() error {
	return dara.Validate(s)
}
