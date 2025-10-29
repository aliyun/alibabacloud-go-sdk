// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpApiOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHttpApiOperationResponseBody
	GetCode() *string
	SetData(v *HttpApiOperationInfo) *GetHttpApiOperationResponseBody
	GetData() *HttpApiOperationInfo
	SetMessage(v string) *GetHttpApiOperationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHttpApiOperationResponseBody
	GetRequestId() *string
}

type GetHttpApiOperationResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Operation information.
	Data *HttpApiOperationInfo `json:"data,omitempty" xml:"data,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// B725275B-50C6-5A49-A9FD-F0332FCB3351
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetHttpApiOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHttpApiOperationResponseBody) GoString() string {
	return s.String()
}

func (s *GetHttpApiOperationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHttpApiOperationResponseBody) GetData() *HttpApiOperationInfo {
	return s.Data
}

func (s *GetHttpApiOperationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHttpApiOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHttpApiOperationResponseBody) SetCode(v string) *GetHttpApiOperationResponseBody {
	s.Code = &v
	return s
}

func (s *GetHttpApiOperationResponseBody) SetData(v *HttpApiOperationInfo) *GetHttpApiOperationResponseBody {
	s.Data = v
	return s
}

func (s *GetHttpApiOperationResponseBody) SetMessage(v string) *GetHttpApiOperationResponseBody {
	s.Message = &v
	return s
}

func (s *GetHttpApiOperationResponseBody) SetRequestId(v string) *GetHttpApiOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHttpApiOperationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
