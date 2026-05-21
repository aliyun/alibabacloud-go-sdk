// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddConnectorResponseBody
	GetCode() *string
	SetData(v *AddConnectorResponseBodyData) *AddConnectorResponseBody
	GetData() *AddConnectorResponseBodyData
	SetMessage(v string) *AddConnectorResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddConnectorResponseBody
	GetRequestId() *string
	SetStatus(v string) *AddConnectorResponseBody
	GetStatus() *string
	SetSuccess(v bool) *AddConnectorResponseBody
	GetSuccess() *bool
}

type AddConnectorResponseBody struct {
	// example:
	//
	// Success
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddConnectorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Current file status does not support delete.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 35A267BF-FBFA-54DB-8394-AA3B0742D833
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *AddConnectorResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddConnectorResponseBody) GetData() *AddConnectorResponseBodyData {
	return s.Data
}

func (s *AddConnectorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddConnectorResponseBody) GetStatus() *string {
	return s.Status
}

func (s *AddConnectorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddConnectorResponseBody) SetCode(v string) *AddConnectorResponseBody {
	s.Code = &v
	return s
}

func (s *AddConnectorResponseBody) SetData(v *AddConnectorResponseBodyData) *AddConnectorResponseBody {
	s.Data = v
	return s
}

func (s *AddConnectorResponseBody) SetMessage(v string) *AddConnectorResponseBody {
	s.Message = &v
	return s
}

func (s *AddConnectorResponseBody) SetRequestId(v string) *AddConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddConnectorResponseBody) SetStatus(v string) *AddConnectorResponseBody {
	s.Status = &v
	return s
}

func (s *AddConnectorResponseBody) SetSuccess(v bool) *AddConnectorResponseBody {
	s.Success = &v
	return s
}

func (s *AddConnectorResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddConnectorResponseBodyData struct {
	// example:
	//
	// conn_file_e0c9db4030b2465a9478028f7d76cd92_1234
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
}

func (s AddConnectorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddConnectorResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddConnectorResponseBodyData) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *AddConnectorResponseBodyData) SetConnectorId(v string) *AddConnectorResponseBodyData {
	s.ConnectorId = &v
	return s
}

func (s *AddConnectorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
