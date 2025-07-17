// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvironmentFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteEnvironmentFeatureResponseBody
	GetCode() *int32
	SetData(v string) *DeleteEnvironmentFeatureResponseBody
	GetData() *string
	SetMessage(v string) *DeleteEnvironmentFeatureResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEnvironmentFeatureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteEnvironmentFeatureResponseBody
	GetSuccess() *bool
}

type DeleteEnvironmentFeatureResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
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
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2A0CEDF1-06FE-44AC-8E21-21A5BE65****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEnvironmentFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvironmentFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentFeatureResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteEnvironmentFeatureResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteEnvironmentFeatureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEnvironmentFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEnvironmentFeatureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteEnvironmentFeatureResponseBody) SetCode(v int32) *DeleteEnvironmentFeatureResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEnvironmentFeatureResponseBody) SetData(v string) *DeleteEnvironmentFeatureResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteEnvironmentFeatureResponseBody) SetMessage(v string) *DeleteEnvironmentFeatureResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEnvironmentFeatureResponseBody) SetRequestId(v string) *DeleteEnvironmentFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnvironmentFeatureResponseBody) SetSuccess(v bool) *DeleteEnvironmentFeatureResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEnvironmentFeatureResponseBody) Validate() error {
	return dara.Validate(s)
}
