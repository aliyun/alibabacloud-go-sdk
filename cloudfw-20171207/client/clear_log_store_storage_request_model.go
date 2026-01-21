// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearLogStoreStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSite(v string) *ClearLogStoreStorageRequest
	GetSite() *string
}

type ClearLogStoreStorageRequest struct {
	// example:
	//
	// cn
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
}

func (s ClearLogStoreStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearLogStoreStorageRequest) GoString() string {
	return s.String()
}

func (s *ClearLogStoreStorageRequest) GetSite() *string {
	return s.Site
}

func (s *ClearLogStoreStorageRequest) SetSite(v string) *ClearLogStoreStorageRequest {
	s.Site = &v
	return s
}

func (s *ClearLogStoreStorageRequest) Validate() error {
	return dara.Validate(s)
}
