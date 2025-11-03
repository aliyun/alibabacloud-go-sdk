// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateConnectionResponseBody
	GetCode() *string
	SetData(v *CreateConnectionResponseBodyData) *CreateConnectionResponseBody
	GetData() *CreateConnectionResponseBodyData
	SetMessage(v string) *CreateConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateConnectionResponseBody
	GetRequestId() *string
}

type CreateConnectionResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message. If the request is successful, success is returned. If the request failed, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7DA60DED-CD36-5837-B848-C01A23D2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConnectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateConnectionResponseBody) GetData() *CreateConnectionResponseBodyData {
	return s.Data
}

func (s *CreateConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConnectionResponseBody) SetCode(v string) *CreateConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConnectionResponseBody) SetData(v *CreateConnectionResponseBodyData) *CreateConnectionResponseBody {
	s.Data = v
	return s
}

func (s *CreateConnectionResponseBody) SetMessage(v string) *CreateConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConnectionResponseBody) SetRequestId(v string) *CreateConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConnectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateConnectionResponseBodyData struct {
	// The connection name.
	//
	// example:
	//
	// connection-demo
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
}

func (s CreateConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateConnectionResponseBodyData) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *CreateConnectionResponseBodyData) SetConnectionName(v string) *CreateConnectionResponseBodyData {
	s.ConnectionName = &v
	return s
}

func (s *CreateConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
