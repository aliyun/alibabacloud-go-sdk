// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerPlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputParam(v string) *TriggerPlaybookRequest
	GetInputParam() *string
	SetPlaybookUuid(v string) *TriggerPlaybookRequest
	GetPlaybookUuid() *string
}

type TriggerPlaybookRequest struct {
	// The input parameters of the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "input1": "xx.xx.xx.xx",
	//
	//     "input2": "7d"
	//
	// }
	InputParam *string `json:"InputParam,omitempty" xml:"InputParam,omitempty"`
	// The playbook UUID.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2a687089-d4dd-47d4-9709-xxxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s TriggerPlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerPlaybookRequest) GoString() string {
	return s.String()
}

func (s *TriggerPlaybookRequest) GetInputParam() *string {
	return s.InputParam
}

func (s *TriggerPlaybookRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *TriggerPlaybookRequest) SetInputParam(v string) *TriggerPlaybookRequest {
	s.InputParam = &v
	return s
}

func (s *TriggerPlaybookRequest) SetPlaybookUuid(v string) *TriggerPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *TriggerPlaybookRequest) Validate() error {
	return dara.Validate(s)
}
