// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarRecordInOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionUuid(v string) *DescribeSoarRecordInOutputRequest
	GetActionUuid() *string
	SetLang(v string) *DescribeSoarRecordInOutputRequest
	GetLang() *string
}

type DescribeSoarRecordInOutputRequest struct {
	// The UUID of the component action.
	//
	// >  You can call the [DescribeSoarTaskAndActions](~~DescribeSoarTaskAndActions~~) operation to query the UUIDs of component actions.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0531ff66-dd05-4f24-84bf-xxxxxxxx
	ActionUuid *string `json:"ActionUuid,omitempty" xml:"ActionUuid,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeSoarRecordInOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarRecordInOutputRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordInOutputRequest) GetActionUuid() *string {
	return s.ActionUuid
}

func (s *DescribeSoarRecordInOutputRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSoarRecordInOutputRequest) SetActionUuid(v string) *DescribeSoarRecordInOutputRequest {
	s.ActionUuid = &v
	return s
}

func (s *DescribeSoarRecordInOutputRequest) SetLang(v string) *DescribeSoarRecordInOutputRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSoarRecordInOutputRequest) Validate() error {
	return dara.Validate(s)
}
