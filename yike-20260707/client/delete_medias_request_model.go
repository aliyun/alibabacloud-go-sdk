// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletePhysicalFiles(v bool) *DeleteMediasRequest
	GetDeletePhysicalFiles() *bool
	SetInputURLs(v string) *DeleteMediasRequest
	GetInputURLs() *string
	SetMediaIds(v string) *DeleteMediasRequest
	GetMediaIds() *string
}

type DeleteMediasRequest struct {
	// Specifies whether to delete the physical files at the same time.
	//
	// example:
	//
	// false
	DeletePhysicalFiles *bool `json:"DeletePhysicalFiles,omitempty" xml:"DeletePhysicalFiles,omitempty"`
	// Not supported.
	InputURLs *string `json:"InputURLs,omitempty" xml:"InputURLs,omitempty"`
	// The media asset IDs, separated by commas. Invalid IDs are added to the IgnoredList.
	//
	// example:
	//
	// ******b48fb04483915d4f2cd8******,******c48fb37407365d4f2cd8******
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
}

func (s DeleteMediasRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediasRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediasRequest) GetDeletePhysicalFiles() *bool {
	return s.DeletePhysicalFiles
}

func (s *DeleteMediasRequest) GetInputURLs() *string {
	return s.InputURLs
}

func (s *DeleteMediasRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *DeleteMediasRequest) SetDeletePhysicalFiles(v bool) *DeleteMediasRequest {
	s.DeletePhysicalFiles = &v
	return s
}

func (s *DeleteMediasRequest) SetInputURLs(v string) *DeleteMediasRequest {
	s.InputURLs = &v
	return s
}

func (s *DeleteMediasRequest) SetMediaIds(v string) *DeleteMediasRequest {
	s.MediaIds = &v
	return s
}

func (s *DeleteMediasRequest) Validate() error {
	return dara.Validate(s)
}
