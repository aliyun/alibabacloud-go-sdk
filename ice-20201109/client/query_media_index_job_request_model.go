// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaIndexJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *QueryMediaIndexJobRequest
	GetMediaId() *string
	SetSearchLibName(v string) *QueryMediaIndexJobRequest
	GetSearchLibName() *string
}

type QueryMediaIndexJobRequest struct {
	// The ID of the media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// c2e77390f75271ec802f0674a2ce6***
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The name of the search library. Default value: ims-default-search-lib.
	//
	// example:
	//
	// test1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
}

func (s QueryMediaIndexJobRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaIndexJobRequest) GoString() string {
	return s.String()
}

func (s *QueryMediaIndexJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *QueryMediaIndexJobRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *QueryMediaIndexJobRequest) SetMediaId(v string) *QueryMediaIndexJobRequest {
	s.MediaId = &v
	return s
}

func (s *QueryMediaIndexJobRequest) SetSearchLibName(v string) *QueryMediaIndexJobRequest {
	s.SearchLibName = &v
	return s
}

func (s *QueryMediaIndexJobRequest) Validate() error {
	return dara.Validate(s)
}
