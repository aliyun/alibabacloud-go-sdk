// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddManifestEditTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddManifestEditTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *AddManifestEditTemplateResponseBody
	GetTemplateId() *string
}

type AddManifestEditTemplateResponseBody struct {
	// **The request ID.**
	//
	// example:
	//
	// req-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// **The ID of the new template.**
	//
	// example:
	//
	// met-abc123
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s AddManifestEditTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddManifestEditTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddManifestEditTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddManifestEditTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AddManifestEditTemplateResponseBody) SetRequestId(v string) *AddManifestEditTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddManifestEditTemplateResponseBody) SetTemplateId(v string) *AddManifestEditTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *AddManifestEditTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
