// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCrTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v []*DescribeCrTemplatesResponseBodyTemplates) *DescribeCrTemplatesResponseBody
	GetTemplates() []*DescribeCrTemplatesResponseBodyTemplates
}

type DescribeCrTemplatesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 74E97AE2-2900-55C1-A069-C3C1EA*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The common YAML templates for the specified type of Istio resource.
	Templates []*DescribeCrTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s DescribeCrTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCrTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCrTemplatesResponseBody) GetTemplates() []*DescribeCrTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *DescribeCrTemplatesResponseBody) SetRequestId(v string) *DescribeCrTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCrTemplatesResponseBody) SetTemplates(v []*DescribeCrTemplatesResponseBodyTemplates) *DescribeCrTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeCrTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCrTemplatesResponseBodyTemplates struct {
	// The Chinese name of the YAML template.
	ChineseName *string `json:"ChineseName,omitempty" xml:"ChineseName,omitempty"`
	// The English name of the YAML template.
	//
	// example:
	//
	// HTTP basic routing
	EnglishName *string `json:"EnglishName,omitempty" xml:"EnglishName,omitempty"`
	// The content in the YAML template.
	//
	// example:
	//
	// apiVersion: networking.istio.io/v1beta1\\nkind: VirtualService\\nmetadata:\\n  name: reviews-route # Name for this VirtualService.\\nspec:\\n  hosts:\\n  - reviews.prod.svc.cluster.local # Service that this VirtualSerivce belongs to. \\n  http:\\n  - name: \\"reviews-route\\" # Name for the route.\\n    route:\\n    - destination: # Uniquely identifies the instances of a service to which all traffic should be forwarded to.\\n        host: reviews.prod.svc.cluster.local # The name of a service from the service registry or ServiceEntry.\\n        port:\\n          number: 8080"
	Yaml *string `json:"Yaml,omitempty" xml:"Yaml,omitempty"`
}

func (s DescribeCrTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeCrTemplatesResponseBodyTemplates) GetChineseName() *string {
	return s.ChineseName
}

func (s *DescribeCrTemplatesResponseBodyTemplates) GetEnglishName() *string {
	return s.EnglishName
}

func (s *DescribeCrTemplatesResponseBodyTemplates) GetYaml() *string {
	return s.Yaml
}

func (s *DescribeCrTemplatesResponseBodyTemplates) SetChineseName(v string) *DescribeCrTemplatesResponseBodyTemplates {
	s.ChineseName = &v
	return s
}

func (s *DescribeCrTemplatesResponseBodyTemplates) SetEnglishName(v string) *DescribeCrTemplatesResponseBodyTemplates {
	s.EnglishName = &v
	return s
}

func (s *DescribeCrTemplatesResponseBodyTemplates) SetYaml(v string) *DescribeCrTemplatesResponseBodyTemplates {
	s.Yaml = &v
	return s
}

func (s *DescribeCrTemplatesResponseBodyTemplates) Validate() error {
	return dara.Validate(s)
}
