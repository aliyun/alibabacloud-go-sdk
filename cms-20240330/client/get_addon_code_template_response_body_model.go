// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddonCodeTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodes(v []*GetAddonCodeTemplateResponseBodyCodes) *GetAddonCodeTemplateResponseBody
	GetCodes() []*GetAddonCodeTemplateResponseBodyCodes
	SetRequestId(v string) *GetAddonCodeTemplateResponseBody
	GetRequestId() *string
}

type GetAddonCodeTemplateResponseBody struct {
	Codes []*GetAddonCodeTemplateResponseBodyCodes `json:"codes,omitempty" xml:"codes,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 0CEC5375-C554-562B-A65F-***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAddonCodeTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAddonCodeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddonCodeTemplateResponseBody) GetCodes() []*GetAddonCodeTemplateResponseBodyCodes {
	return s.Codes
}

func (s *GetAddonCodeTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAddonCodeTemplateResponseBody) SetCodes(v []*GetAddonCodeTemplateResponseBodyCodes) *GetAddonCodeTemplateResponseBody {
	s.Codes = v
	return s
}

func (s *GetAddonCodeTemplateResponseBody) SetRequestId(v string) *GetAddonCodeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAddonCodeTemplateResponseBody) Validate() error {
	if s.Codes != nil {
		for _, item := range s.Codes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAddonCodeTemplateResponseBodyCodes struct {
	// example:
	//
	// javascript\\nnpm install @arms/rum-miniapp --save
	CodeTemplate *string `json:"codeTemplate,omitempty" xml:"codeTemplate,omitempty"`
	// example:
	//
	// cs-default-CS-cs-default-1753236205394-cs-default-CS-kubeApiserver
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAddonCodeTemplateResponseBodyCodes) String() string {
	return dara.Prettify(s)
}

func (s GetAddonCodeTemplateResponseBodyCodes) GoString() string {
	return s.String()
}

func (s *GetAddonCodeTemplateResponseBodyCodes) GetCodeTemplate() *string {
	return s.CodeTemplate
}

func (s *GetAddonCodeTemplateResponseBodyCodes) GetName() *string {
	return s.Name
}

func (s *GetAddonCodeTemplateResponseBodyCodes) SetCodeTemplate(v string) *GetAddonCodeTemplateResponseBodyCodes {
	s.CodeTemplate = &v
	return s
}

func (s *GetAddonCodeTemplateResponseBodyCodes) SetName(v string) *GetAddonCodeTemplateResponseBodyCodes {
	s.Name = &v
	return s
}

func (s *GetAddonCodeTemplateResponseBodyCodes) Validate() error {
	return dara.Validate(s)
}
