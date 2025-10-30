// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafStartStepsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSafStartStepsResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeSafStartStepsResponseBodyResultObject) *DescribeSafStartStepsResponseBody
	GetResultObject() []*DescribeSafStartStepsResponseBodyResultObject
}

type DescribeSafStartStepsResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*DescribeSafStartStepsResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeSafStartStepsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafStartStepsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSafStartStepsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSafStartStepsResponseBody) GetResultObject() []*DescribeSafStartStepsResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSafStartStepsResponseBody) SetRequestId(v string) *DescribeSafStartStepsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSafStartStepsResponseBody) SetResultObject(v []*DescribeSafStartStepsResponseBodyResultObject) *DescribeSafStartStepsResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSafStartStepsResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSafStartStepsResponseBodyResultObject struct {
	// Step content
	//
	// example:
	//
	// https://help.aliyun.com/document_detail/177689.html
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Step title
	//
	// example:
	//
	// ios
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Type
	//
	// example:
	//
	// URL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeSafStartStepsResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafStartStepsResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSafStartStepsResponseBodyResultObject) GetContent() *string {
	return s.Content
}

func (s *DescribeSafStartStepsResponseBodyResultObject) GetId() *string {
	return s.Id
}

func (s *DescribeSafStartStepsResponseBodyResultObject) GetType() *string {
	return s.Type
}

func (s *DescribeSafStartStepsResponseBodyResultObject) SetContent(v string) *DescribeSafStartStepsResponseBodyResultObject {
	s.Content = &v
	return s
}

func (s *DescribeSafStartStepsResponseBodyResultObject) SetId(v string) *DescribeSafStartStepsResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeSafStartStepsResponseBodyResultObject) SetType(v string) *DescribeSafStartStepsResponseBodyResultObject {
	s.Type = &v
	return s
}

func (s *DescribeSafStartStepsResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
