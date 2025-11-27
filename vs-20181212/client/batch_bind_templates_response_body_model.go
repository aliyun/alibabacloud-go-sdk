// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchBindTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchBindTemplatesResponseBody
	GetRequestId() *string
	SetResults(v []*BatchBindTemplatesResponseBodyResults) *BatchBindTemplatesResponseBody
	GetResults() []*BatchBindTemplatesResponseBodyResults
}

type BatchBindTemplatesResponseBody struct {
	// example:
	//
	// 20D0DFCE-5DB7-5D83-BD82-8482F2327636
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchBindTemplatesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchBindTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchBindTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchBindTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchBindTemplatesResponseBody) GetResults() []*BatchBindTemplatesResponseBodyResults {
	return s.Results
}

func (s *BatchBindTemplatesResponseBody) SetRequestId(v string) *BatchBindTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchBindTemplatesResponseBody) SetResults(v []*BatchBindTemplatesResponseBodyResults) *BatchBindTemplatesResponseBody {
	s.Results = v
	return s
}

func (s *BatchBindTemplatesResponseBody) Validate() error {
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

type BatchBindTemplatesResponseBodyResults struct {
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
}

func (s BatchBindTemplatesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchBindTemplatesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchBindTemplatesResponseBodyResults) GetError() *string {
	return s.Error
}

func (s *BatchBindTemplatesResponseBodyResults) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BatchBindTemplatesResponseBodyResults) GetInstanceType() *string {
	return s.InstanceType
}

func (s *BatchBindTemplatesResponseBodyResults) GetTemplateId() *string {
	return s.TemplateId
}

func (s *BatchBindTemplatesResponseBodyResults) SetError(v string) *BatchBindTemplatesResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchBindTemplatesResponseBodyResults) SetInstanceId(v string) *BatchBindTemplatesResponseBodyResults {
	s.InstanceId = &v
	return s
}

func (s *BatchBindTemplatesResponseBodyResults) SetInstanceType(v string) *BatchBindTemplatesResponseBodyResults {
	s.InstanceType = &v
	return s
}

func (s *BatchBindTemplatesResponseBodyResults) SetTemplateId(v string) *BatchBindTemplatesResponseBodyResults {
	s.TemplateId = &v
	return s
}

func (s *BatchBindTemplatesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
