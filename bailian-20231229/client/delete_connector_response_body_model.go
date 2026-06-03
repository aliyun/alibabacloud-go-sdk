// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteConnectorResponseBody
	GetCode() *string
	SetData(v *DeleteConnectorResponseBodyData) *DeleteConnectorResponseBody
	GetData() *DeleteConnectorResponseBodyData
	SetMessage(v string) *DeleteConnectorResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteConnectorResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteConnectorResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DeleteConnectorResponseBody
	GetSuccess() *bool
}

type DeleteConnectorResponseBody struct {
	// example:
	//
	// Index.InvalidParameter
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteConnectorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// User not authorized to operate on the specified resource
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1C139002-0EC5-584C-A755-4B8B9FA080BE
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

func (s DeleteConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConnectorResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteConnectorResponseBody) GetData() *DeleteConnectorResponseBodyData {
	return s.Data
}

func (s *DeleteConnectorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConnectorResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteConnectorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteConnectorResponseBody) SetCode(v string) *DeleteConnectorResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConnectorResponseBody) SetData(v *DeleteConnectorResponseBodyData) *DeleteConnectorResponseBody {
	s.Data = v
	return s
}

func (s *DeleteConnectorResponseBody) SetMessage(v string) *DeleteConnectorResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConnectorResponseBody) SetRequestId(v string) *DeleteConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConnectorResponseBody) SetStatus(v string) *DeleteConnectorResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteConnectorResponseBody) SetSuccess(v bool) *DeleteConnectorResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteConnectorResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteConnectorResponseBodyData struct {
	// example:
	//
	// conn_file_e0c9db4030b2465a9478028f7d76cd92_1234
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
}

func (s DeleteConnectorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectorResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteConnectorResponseBodyData) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *DeleteConnectorResponseBodyData) SetConnectorId(v string) *DeleteConnectorResponseBodyData {
	s.ConnectorId = &v
	return s
}

func (s *DeleteConnectorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
