// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetConnectorResponseBody
	GetCode() *string
	SetData(v *GetConnectorResponseBodyData) *GetConnectorResponseBody
	GetData() *GetConnectorResponseBodyData
	SetMessage(v string) *GetConnectorResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConnectorResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetConnectorResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetConnectorResponseBody
	GetSuccess() *bool
}

type GetConnectorResponseBody struct {
	// example:
	//
	// success
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetConnectorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Requests throttling triggered.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7BA8ADD9-53D6-53F0-918F-A1E776AD230E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectorResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConnectorResponseBody) GetData() *GetConnectorResponseBodyData {
	return s.Data
}

func (s *GetConnectorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConnectorResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetConnectorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConnectorResponseBody) SetCode(v string) *GetConnectorResponseBody {
	s.Code = &v
	return s
}

func (s *GetConnectorResponseBody) SetData(v *GetConnectorResponseBodyData) *GetConnectorResponseBody {
	s.Data = v
	return s
}

func (s *GetConnectorResponseBody) SetMessage(v string) *GetConnectorResponseBody {
	s.Message = &v
	return s
}

func (s *GetConnectorResponseBody) SetRequestId(v string) *GetConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectorResponseBody) SetStatus(v string) *GetConnectorResponseBody {
	s.Status = &v
	return s
}

func (s *GetConnectorResponseBody) SetSuccess(v bool) *GetConnectorResponseBody {
	s.Success = &v
	return s
}

func (s *GetConnectorResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConnectorResponseBodyData struct {
	// example:
	//
	// conn_file_e0c9db4030b2465a9478028f7d76cd92_1234
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// example:
	//
	// name
	ConnectorName *string `json:"ConnectorName,omitempty" xml:"ConnectorName,omitempty"`
	// example:
	//
	// FILE
	ConnectorType *string `json:"ConnectorType,omitempty" xml:"ConnectorType,omitempty"`
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetConnectorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetConnectorResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConnectorResponseBodyData) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *GetConnectorResponseBodyData) GetConnectorName() *string {
	return s.ConnectorName
}

func (s *GetConnectorResponseBodyData) GetConnectorType() *string {
	return s.ConnectorType
}

func (s *GetConnectorResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetConnectorResponseBodyData) SetConnectorId(v string) *GetConnectorResponseBodyData {
	s.ConnectorId = &v
	return s
}

func (s *GetConnectorResponseBodyData) SetConnectorName(v string) *GetConnectorResponseBodyData {
	s.ConnectorName = &v
	return s
}

func (s *GetConnectorResponseBodyData) SetConnectorType(v string) *GetConnectorResponseBodyData {
	s.ConnectorType = &v
	return s
}

func (s *GetConnectorResponseBodyData) SetDescription(v string) *GetConnectorResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetConnectorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
