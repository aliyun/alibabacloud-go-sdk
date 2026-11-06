// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticVpcEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteAgenticVpcEndpointResponseBodyData) *DeleteAgenticVpcEndpointResponseBody
	GetData() *DeleteAgenticVpcEndpointResponseBodyData
	SetErrorCode(v string) *DeleteAgenticVpcEndpointResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteAgenticVpcEndpointResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteAgenticVpcEndpointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAgenticVpcEndpointResponseBody
	GetSuccess() *bool
}

type DeleteAgenticVpcEndpointResponseBody struct {
	// The deletion acceptance result. For duplicate requests, downstream errors are retained and idempotence is not guaranteed.
	Data *DeleteAgenticVpcEndpointResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when a request exception occurs.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAgenticVpcEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgenticVpcEndpointResponseBody) GetData() *DeleteAgenticVpcEndpointResponseBodyData {
	return s.Data
}

func (s *DeleteAgenticVpcEndpointResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteAgenticVpcEndpointResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteAgenticVpcEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgenticVpcEndpointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAgenticVpcEndpointResponseBody) SetData(v *DeleteAgenticVpcEndpointResponseBodyData) *DeleteAgenticVpcEndpointResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAgenticVpcEndpointResponseBody) SetErrorCode(v string) *DeleteAgenticVpcEndpointResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAgenticVpcEndpointResponseBody) SetErrorMessage(v string) *DeleteAgenticVpcEndpointResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteAgenticVpcEndpointResponseBody) SetRequestId(v string) *DeleteAgenticVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgenticVpcEndpointResponseBody) SetSuccess(v bool) *DeleteAgenticVpcEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAgenticVpcEndpointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAgenticVpcEndpointResponseBodyData struct {
	// The ID of the endpoint accepted for deletion.
	//
	// example:
	//
	// ep-7xvra4ad6619266ba43d
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The deletion status. A value of DELETING only indicates that the asynchronous deletion request is accepted, not that the deletion in the cloud is complete.
	//
	// example:
	//
	// RELEASED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteAgenticVpcEndpointResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticVpcEndpointResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteAgenticVpcEndpointResponseBodyData) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DeleteAgenticVpcEndpointResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DeleteAgenticVpcEndpointResponseBodyData) SetEndpointId(v string) *DeleteAgenticVpcEndpointResponseBodyData {
	s.EndpointId = &v
	return s
}

func (s *DeleteAgenticVpcEndpointResponseBodyData) SetStatus(v string) *DeleteAgenticVpcEndpointResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteAgenticVpcEndpointResponseBodyData) Validate() error {
	return dara.Validate(s)
}
