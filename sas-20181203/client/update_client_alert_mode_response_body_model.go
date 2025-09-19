// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientAlertModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateClientAlertModeResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateClientAlertModeResponseBody
	GetRequestId() *string
}

type UpdateClientAlertModeResponseBody struct {
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 259E3E77-CA6D-5407-84A5-3A1C98D12F14
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClientAlertModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientAlertModeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClientAlertModeResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateClientAlertModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClientAlertModeResponseBody) SetData(v bool) *UpdateClientAlertModeResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateClientAlertModeResponseBody) SetRequestId(v string) *UpdateClientAlertModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClientAlertModeResponseBody) Validate() error {
	return dara.Validate(s)
}
