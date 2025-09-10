// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaTemplateServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceDirectoryId(v string) *GetQuotaTemplateServiceStatusRequest
	GetResourceDirectoryId() *string
}

type GetQuotaTemplateServiceStatusRequest struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-pG****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
}

func (s GetQuotaTemplateServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaTemplateServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaTemplateServiceStatusRequest) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *GetQuotaTemplateServiceStatusRequest) SetResourceDirectoryId(v string) *GetQuotaTemplateServiceStatusRequest {
	s.ResourceDirectoryId = &v
	return s
}

func (s *GetQuotaTemplateServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
