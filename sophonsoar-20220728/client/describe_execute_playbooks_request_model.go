// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExecutePlaybooksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputMode(v string) *DescribeExecutePlaybooksRequest
	GetInputMode() *string
	SetLang(v string) *DescribeExecutePlaybooksRequest
	GetLang() *string
	SetParamType(v string) *DescribeExecutePlaybooksRequest
	GetParamType() *string
	SetPlaybookName(v string) *DescribeExecutePlaybooksRequest
	GetPlaybookName() *string
	SetUuid(v string) *DescribeExecutePlaybooksRequest
	GetUuid() *string
}

type DescribeExecutePlaybooksRequest struct {
	// The entity type of the input parameter for the playbook. To query multiple entity types, separate them with commas.
	//
	// - **ip**: IP entity.
	//
	// - **file**: file entity.
	//
	// - **process**: process entity.
	//
	// - **incident**: event entity.
	//
	// example:
	//
	// ip,file,process,host
	InputMode *string `json:"InputMode,omitempty" xml:"InputMode,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of the input parameter for the playbook.
	//
	// - **template-ip**: IP request template.
	//
	// - **template-file**: file request template.
	//
	// - **template-process**: process request template.
	//
	// - **custom**: custom parameter.
	//
	// example:
	//
	// custom
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The name of the playbook. Fuzzy search is supported.
	//
	// example:
	//
	// demo_test
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	// The UUID of the playbook.
	//
	// > Call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to obtain this parameter.
	//
	// example:
	//
	// f916b93e-e814-459f-9662-xxxxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeExecutePlaybooksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutePlaybooksRequest) GoString() string {
	return s.String()
}

func (s *DescribeExecutePlaybooksRequest) GetInputMode() *string {
	return s.InputMode
}

func (s *DescribeExecutePlaybooksRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeExecutePlaybooksRequest) GetParamType() *string {
	return s.ParamType
}

func (s *DescribeExecutePlaybooksRequest) GetPlaybookName() *string {
	return s.PlaybookName
}

func (s *DescribeExecutePlaybooksRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeExecutePlaybooksRequest) SetInputMode(v string) *DescribeExecutePlaybooksRequest {
	s.InputMode = &v
	return s
}

func (s *DescribeExecutePlaybooksRequest) SetLang(v string) *DescribeExecutePlaybooksRequest {
	s.Lang = &v
	return s
}

func (s *DescribeExecutePlaybooksRequest) SetParamType(v string) *DescribeExecutePlaybooksRequest {
	s.ParamType = &v
	return s
}

func (s *DescribeExecutePlaybooksRequest) SetPlaybookName(v string) *DescribeExecutePlaybooksRequest {
	s.PlaybookName = &v
	return s
}

func (s *DescribeExecutePlaybooksRequest) SetUuid(v string) *DescribeExecutePlaybooksRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeExecutePlaybooksRequest) Validate() error {
	return dara.Validate(s)
}
