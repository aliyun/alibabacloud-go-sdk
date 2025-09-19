// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupAutoConfigStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetBackupAutoConfigStatusResponseBodyData) *GetBackupAutoConfigStatusResponseBody
	GetData() *GetBackupAutoConfigStatusResponseBodyData
	SetRequestId(v string) *GetBackupAutoConfigStatusResponseBody
	GetRequestId() *string
}

type GetBackupAutoConfigStatusResponseBody struct {
	// The response parameters.
	Data *GetBackupAutoConfigStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 898F7AA7-CECD-5EC7-AF4D-664C601B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBackupAutoConfigStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBackupAutoConfigStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetBackupAutoConfigStatusResponseBody) GetData() *GetBackupAutoConfigStatusResponseBodyData {
	return s.Data
}

func (s *GetBackupAutoConfigStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBackupAutoConfigStatusResponseBody) SetData(v *GetBackupAutoConfigStatusResponseBodyData) *GetBackupAutoConfigStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetBackupAutoConfigStatusResponseBody) SetRequestId(v string) *GetBackupAutoConfigStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBackupAutoConfigStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBackupAutoConfigStatusResponseBodyData struct {
	// Indicates whether the anti-ransomware policy for servers can be automatically configured by the managed anti-ransomware feature. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// false
	CanConfigAuto *bool `json:"CanConfigAuto,omitempty" xml:"CanConfigAuto,omitempty"`
}

func (s GetBackupAutoConfigStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBackupAutoConfigStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBackupAutoConfigStatusResponseBodyData) GetCanConfigAuto() *bool {
	return s.CanConfigAuto
}

func (s *GetBackupAutoConfigStatusResponseBodyData) SetCanConfigAuto(v bool) *GetBackupAutoConfigStatusResponseBodyData {
	s.CanConfigAuto = &v
	return s
}

func (s *GetBackupAutoConfigStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
