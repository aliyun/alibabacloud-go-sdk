// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoutineRelatedDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *DescribeRoutineRelatedDomainsResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *DescribeRoutineRelatedDomainsResponseBody
	GetRequestId() *string
}

type DescribeRoutineRelatedDomainsResponseBody struct {
	// The domain names associated with a routine.
	//
	// example:
	//
	// "Domains": [
	//
	//             "xxx.com",
	//
	//             "yyy.com",
	//
	//             ...
	//
	//         ]
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FC0E34AC-0239-44A7-AB0E-800DE522C8DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRoutineRelatedDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoutineRelatedDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoutineRelatedDomainsResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *DescribeRoutineRelatedDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRoutineRelatedDomainsResponseBody) SetContent(v map[string]interface{}) *DescribeRoutineRelatedDomainsResponseBody {
	s.Content = v
	return s
}

func (s *DescribeRoutineRelatedDomainsResponseBody) SetRequestId(v string) *DescribeRoutineRelatedDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRoutineRelatedDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}
