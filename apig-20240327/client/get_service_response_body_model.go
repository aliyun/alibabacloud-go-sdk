// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetServiceResponseBody
	GetCode() *string
	SetData(v *Service) *GetServiceResponseBody
	GetData() *Service
	SetMessage(v string) *GetServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetServiceResponseBody
	GetRequestId() *string
}

type GetServiceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The service details.
	Data *Service `json:"data,omitempty" xml:"data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8FA9BB94-915B-5299-A694-49FCC7F5DD00
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetServiceResponseBody) GetData() *Service {
	return s.Data
}

func (s *GetServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceResponseBody) SetCode(v string) *GetServiceResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceResponseBody) SetData(v *Service) *GetServiceResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceResponseBody) SetMessage(v string) *GetServiceResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceResponseBody) SetRequestId(v string) *GetServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
