// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTemplatesResponseBody
	GetRequestId() *string
	SetResult(v []*DescribeTemplatesResponseBodyResult) *DescribeTemplatesResponseBody
	GetResult() []*DescribeTemplatesResponseBodyResult
}

type DescribeTemplatesResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeTemplatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTemplatesResponseBody) GetResult() []*DescribeTemplatesResponseBodyResult {
	return s.Result
}

func (s *DescribeTemplatesResponseBody) SetRequestId(v string) *DescribeTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetResult(v []*DescribeTemplatesResponseBodyResult) *DescribeTemplatesResponseBody {
	s.Result = v
	return s
}

func (s *DescribeTemplatesResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTemplatesResponseBodyResult struct {
	// example:
	//
	// {\\n\\t\\"persistent\\":{\\n\\t\\t\\"search\\":{\\n\\t\\t\\t\\"max_buckets\\":\\"10000\\"\\n\\t\\t}\\n\\t}\\n}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// dynamicSettings
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s DescribeTemplatesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *DescribeTemplatesResponseBodyResult) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeTemplatesResponseBodyResult) SetContent(v string) *DescribeTemplatesResponseBodyResult {
	s.Content = &v
	return s
}

func (s *DescribeTemplatesResponseBodyResult) SetTemplateName(v string) *DescribeTemplatesResponseBodyResult {
	s.TemplateName = &v
	return s
}

func (s *DescribeTemplatesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
