// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteManifestEditTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteManifestEditTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *DeleteManifestEditTemplateResponseBody
	GetTemplateId() *string
}

type DeleteManifestEditTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the deleted ManifestEdit template.
	//
	// example:
	//
	// 9c6581d59240de5e258ffa4c71a3e01b
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteManifestEditTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteManifestEditTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteManifestEditTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteManifestEditTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteManifestEditTemplateResponseBody) SetRequestId(v string) *DeleteManifestEditTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteManifestEditTemplateResponseBody) SetTemplateId(v string) *DeleteManifestEditTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *DeleteManifestEditTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
