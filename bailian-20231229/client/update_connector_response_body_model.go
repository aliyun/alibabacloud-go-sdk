// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateConnectorResponseBody
	GetCode() *string
	SetData(v *UpdateConnectorResponseBodyData) *UpdateConnectorResponseBody
	GetData() *UpdateConnectorResponseBodyData
	SetMessage(v string) *UpdateConnectorResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateConnectorResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateConnectorResponseBody
	GetStatus() *string
	SetSuccess(v bool) *UpdateConnectorResponseBody
	GetSuccess() *bool
}

type UpdateConnectorResponseBody struct {
	// example:
	//
	// Index.Forbidden
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateConnectorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 778C0B3B-03C1-5FC1-A947-36EDD13606AB
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

func (s UpdateConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnectorResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateConnectorResponseBody) GetData() *UpdateConnectorResponseBodyData {
	return s.Data
}

func (s *UpdateConnectorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConnectorResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateConnectorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateConnectorResponseBody) SetCode(v string) *UpdateConnectorResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConnectorResponseBody) SetData(v *UpdateConnectorResponseBodyData) *UpdateConnectorResponseBody {
	s.Data = v
	return s
}

func (s *UpdateConnectorResponseBody) SetMessage(v string) *UpdateConnectorResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConnectorResponseBody) SetRequestId(v string) *UpdateConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConnectorResponseBody) SetStatus(v string) *UpdateConnectorResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateConnectorResponseBody) SetSuccess(v bool) *UpdateConnectorResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateConnectorResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateConnectorResponseBodyData struct {
	// example:
	//
	// conn_file_e0c9db4030b2465a9478028f7d76cd92_1234
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
}

func (s UpdateConnectorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateConnectorResponseBodyData) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *UpdateConnectorResponseBodyData) SetConnectorId(v string) *UpdateConnectorResponseBodyData {
	s.ConnectorId = &v
	return s
}

func (s *UpdateConnectorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
