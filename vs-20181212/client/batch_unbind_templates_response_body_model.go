// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUnbindTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchUnbindTemplatesResponseBody
	GetRequestId() *string
	SetResults(v []*BatchUnbindTemplatesResponseBodyResults) *BatchUnbindTemplatesResponseBody
	GetResults() []*BatchUnbindTemplatesResponseBodyResults
}

type BatchUnbindTemplatesResponseBody struct {
	// example:
	//
	// 90B377DF-C874-5BBD-B957-42C4C06AFECE
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchUnbindTemplatesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchUnbindTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUnbindTemplatesResponseBody) GetResults() []*BatchUnbindTemplatesResponseBodyResults {
	return s.Results
}

func (s *BatchUnbindTemplatesResponseBody) SetRequestId(v string) *BatchUnbindTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUnbindTemplatesResponseBody) SetResults(v []*BatchUnbindTemplatesResponseBodyResults) *BatchUnbindTemplatesResponseBody {
	s.Results = v
	return s
}

func (s *BatchUnbindTemplatesResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchUnbindTemplatesResponseBodyResults struct {
	// example:
	//
	// some error
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 323*****994-cn-qingdao
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// group
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 323*****998-cn-qingdao
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// snapshot
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s BatchUnbindTemplatesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindTemplatesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplatesResponseBodyResults) GetError() *string {
	return s.Error
}

func (s *BatchUnbindTemplatesResponseBodyResults) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BatchUnbindTemplatesResponseBodyResults) GetInstanceType() *string {
	return s.InstanceType
}

func (s *BatchUnbindTemplatesResponseBodyResults) GetTemplateId() *string {
	return s.TemplateId
}

func (s *BatchUnbindTemplatesResponseBodyResults) GetTemplateType() *string {
	return s.TemplateType
}

func (s *BatchUnbindTemplatesResponseBodyResults) SetError(v string) *BatchUnbindTemplatesResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchUnbindTemplatesResponseBodyResults) SetInstanceId(v string) *BatchUnbindTemplatesResponseBodyResults {
	s.InstanceId = &v
	return s
}

func (s *BatchUnbindTemplatesResponseBodyResults) SetInstanceType(v string) *BatchUnbindTemplatesResponseBodyResults {
	s.InstanceType = &v
	return s
}

func (s *BatchUnbindTemplatesResponseBodyResults) SetTemplateId(v string) *BatchUnbindTemplatesResponseBodyResults {
	s.TemplateId = &v
	return s
}

func (s *BatchUnbindTemplatesResponseBodyResults) SetTemplateType(v string) *BatchUnbindTemplatesResponseBodyResults {
	s.TemplateType = &v
	return s
}

func (s *BatchUnbindTemplatesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
