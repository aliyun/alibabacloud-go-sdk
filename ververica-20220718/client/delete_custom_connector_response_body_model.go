// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*TableMeta) *DeleteCustomConnectorResponseBody
	GetData() []*TableMeta
	SetErrorCode(v string) *DeleteCustomConnectorResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteCustomConnectorResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeleteCustomConnectorResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteCustomConnectorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomConnectorResponseBody
	GetSuccess() *bool
}

type DeleteCustomConnectorResponseBody struct {
	// If the value of success was true, a list of deployments in which custom connectors were deleted was returned. If the value of success was false, a null value was returned.
	Data []*TableMeta `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteCustomConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomConnectorResponseBody) GetData() []*TableMeta {
	return s.Data
}

func (s *DeleteCustomConnectorResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteCustomConnectorResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteCustomConnectorResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteCustomConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomConnectorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomConnectorResponseBody) SetData(v []*TableMeta) *DeleteCustomConnectorResponseBody {
	s.Data = v
	return s
}

func (s *DeleteCustomConnectorResponseBody) SetErrorCode(v string) *DeleteCustomConnectorResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteCustomConnectorResponseBody) SetErrorMessage(v string) *DeleteCustomConnectorResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteCustomConnectorResponseBody) SetHttpCode(v int32) *DeleteCustomConnectorResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteCustomConnectorResponseBody) SetRequestId(v string) *DeleteCustomConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomConnectorResponseBody) SetSuccess(v bool) *DeleteCustomConnectorResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomConnectorResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
