// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSolutionInstanceConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []map[string]interface{}) *DescribeSolutionInstanceConfigurationResponseBody
	GetData() []map[string]interface{}
	SetRequestId(v string) *DescribeSolutionInstanceConfigurationResponseBody
	GetRequestId() *string
}

type DescribeSolutionInstanceConfigurationResponseBody struct {
	// The returned data.
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 11B55F58-D3A4-4A9B-9596-342420D0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSolutionInstanceConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSolutionInstanceConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSolutionInstanceConfigurationResponseBody) GetData() []map[string]interface{} {
	return s.Data
}

func (s *DescribeSolutionInstanceConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSolutionInstanceConfigurationResponseBody) SetData(v []map[string]interface{}) *DescribeSolutionInstanceConfigurationResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSolutionInstanceConfigurationResponseBody) SetRequestId(v string) *DescribeSolutionInstanceConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSolutionInstanceConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
