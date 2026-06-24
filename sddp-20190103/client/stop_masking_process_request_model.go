// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMaskingProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *StopMaskingProcessRequest
	GetId() *int64
	SetLang(v string) *StopMaskingProcessRequest
	GetLang() *string
}

type StopMaskingProcessRequest struct {
	// The unique ID of the data masking task. You can obtain the ID of a data masking task from the return value of the [DescribeDataMaskingTasks](~~DescribeDataMaskingTasks~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh_cn**: Chinese (Simplified). This is the default value.
	//
	// - **en_us**: English (US).
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s StopMaskingProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s StopMaskingProcessRequest) GoString() string {
	return s.String()
}

func (s *StopMaskingProcessRequest) GetId() *int64 {
	return s.Id
}

func (s *StopMaskingProcessRequest) GetLang() *string {
	return s.Lang
}

func (s *StopMaskingProcessRequest) SetId(v int64) *StopMaskingProcessRequest {
	s.Id = &v
	return s
}

func (s *StopMaskingProcessRequest) SetLang(v string) *StopMaskingProcessRequest {
	s.Lang = &v
	return s
}

func (s *StopMaskingProcessRequest) Validate() error {
	return dara.Validate(s)
}
