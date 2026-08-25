// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectoryStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetDirectoryStatisticsRequest
	GetDirectoryId() *string
}

type GetDirectoryStatisticsRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetDirectoryStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetDirectoryStatisticsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetDirectoryStatisticsRequest) SetDirectoryId(v string) *GetDirectoryStatisticsRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetDirectoryStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
