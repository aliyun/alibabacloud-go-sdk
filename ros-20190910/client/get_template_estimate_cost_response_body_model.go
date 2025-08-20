// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateEstimateCostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTemplateEstimateCostResponseBody
	GetRequestId() *string
	SetResources(v map[string]interface{}) *GetTemplateEstimateCostResponseBody
	GetResources() map[string]interface{}
}

type GetTemplateEstimateCostResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6DEA36EF-C97D-5658-A4AC-4F5DB08D1A89
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource details.
	//
	// example:
	//
	// [{\\"LogicalResourceIdPattern\\": \\"wait1\\", \\"ResourceType\\": \\"time_sleep\\", \\"ResourcePath\\": \\"wait1\\"}, {\\"LogicalResourceIdPattern\\": \\"wait2\\", \\"ResourceType\\": \\"time_sleep\\", \\"ResourcePath\\": \\"wait2\\"}]
	Resources map[string]interface{} `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s GetTemplateEstimateCostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateEstimateCostResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateEstimateCostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateEstimateCostResponseBody) GetResources() map[string]interface{} {
	return s.Resources
}

func (s *GetTemplateEstimateCostResponseBody) SetRequestId(v string) *GetTemplateEstimateCostResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateEstimateCostResponseBody) SetResources(v map[string]interface{}) *GetTemplateEstimateCostResponseBody {
	s.Resources = v
	return s
}

func (s *GetTemplateEstimateCostResponseBody) Validate() error {
	return dara.Validate(s)
}
