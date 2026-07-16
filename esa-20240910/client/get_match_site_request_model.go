// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMatchSiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordName(v string) *GetMatchSiteRequest
	GetRecordName() *string
}

type GetMatchSiteRequest struct {
	// The record name.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
}

func (s GetMatchSiteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMatchSiteRequest) GoString() string {
	return s.String()
}

func (s *GetMatchSiteRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *GetMatchSiteRequest) SetRecordName(v string) *GetMatchSiteRequest {
	s.RecordName = &v
	return s
}

func (s *GetMatchSiteRequest) Validate() error {
	return dara.Validate(s)
}
