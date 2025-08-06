// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAttachedMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaIds(v string) *DeleteAttachedMediaRequest
	GetMediaIds() *string
}

type DeleteAttachedMediaRequest struct {
	// The ID of the auxiliary media asset that you want to delete.
	//
	// 	- Separate multiple IDs with commas (,). You can specify up to 20 IDs.
	//
	// 	- You can obtain the ID from the response to the [CreateUploadAttachedMedia](~~CreateUploadAttachedMedia~~) operation that you call to obtain the upload URL and credential.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8bc8e94fe4e55abde85718****,eb186180e989dd56****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
}

func (s DeleteAttachedMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAttachedMediaRequest) GoString() string {
	return s.String()
}

func (s *DeleteAttachedMediaRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *DeleteAttachedMediaRequest) SetMediaIds(v string) *DeleteAttachedMediaRequest {
	s.MediaIds = &v
	return s
}

func (s *DeleteAttachedMediaRequest) Validate() error {
	return dara.Validate(s)
}
