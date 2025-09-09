// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManualTriggerMaskingProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ManualTriggerMaskingProcessRequest
	GetId() *int64
	SetLang(v string) *ManualTriggerMaskingProcessRequest
	GetLang() *string
}

type ManualTriggerMaskingProcessRequest struct {
	// The ID of the de-identification task.
	//
	// The ID of the de-identification task is a string. You can call the DescribeDataMaskingTasks operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response, default value zh_cn. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ManualTriggerMaskingProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s ManualTriggerMaskingProcessRequest) GoString() string {
	return s.String()
}

func (s *ManualTriggerMaskingProcessRequest) GetId() *int64 {
	return s.Id
}

func (s *ManualTriggerMaskingProcessRequest) GetLang() *string {
	return s.Lang
}

func (s *ManualTriggerMaskingProcessRequest) SetId(v int64) *ManualTriggerMaskingProcessRequest {
	s.Id = &v
	return s
}

func (s *ManualTriggerMaskingProcessRequest) SetLang(v string) *ManualTriggerMaskingProcessRequest {
	s.Lang = &v
	return s
}

func (s *ManualTriggerMaskingProcessRequest) Validate() error {
	return dara.Validate(s)
}
