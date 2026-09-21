// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCnameFlatteningRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlattenMode(v string) *UpdateCnameFlatteningRequest
	GetFlattenMode() *string
	SetSiteId(v int64) *UpdateCnameFlatteningRequest
	GetSiteId() *int64
}

type UpdateCnameFlatteningRequest struct {
	// The flattening mode. Valid values:
	//
	// - flatten_all: flattens all records.
	//
	// - flatten_at_root: flattens only the root domain. This is the default value.
	//
	// This parameter is required.
	//
	// example:
	//
	// flatten_all
	FlattenMode *string `json:"FlattenMode,omitempty" xml:"FlattenMode,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateCnameFlatteningRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCnameFlatteningRequest) GoString() string {
	return s.String()
}

func (s *UpdateCnameFlatteningRequest) GetFlattenMode() *string {
	return s.FlattenMode
}

func (s *UpdateCnameFlatteningRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateCnameFlatteningRequest) SetFlattenMode(v string) *UpdateCnameFlatteningRequest {
	s.FlattenMode = &v
	return s
}

func (s *UpdateCnameFlatteningRequest) SetSiteId(v int64) *UpdateCnameFlatteningRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateCnameFlatteningRequest) Validate() error {
	return dara.Validate(s)
}
