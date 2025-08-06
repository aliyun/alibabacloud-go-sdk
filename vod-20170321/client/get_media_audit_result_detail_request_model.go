// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaAuditResultDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *GetMediaAuditResultDetailRequest
	GetMediaId() *string
	SetPageNo(v int32) *GetMediaAuditResultDetailRequest
	GetPageNo() *int32
}

type GetMediaAuditResultDetailRequest struct {
	// The ID of the video.
	//
	// This parameter is required.
	//
	// example:
	//
	// 93ab850b4f6f*****54b6e91d24d81d4
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The page number. The default value is **1**. A maximum of **20*	- records can be returned on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
}

func (s GetMediaAuditResultDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultDetailRequest) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultDetailRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaAuditResultDetailRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetMediaAuditResultDetailRequest) SetMediaId(v string) *GetMediaAuditResultDetailRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaAuditResultDetailRequest) SetPageNo(v int32) *GetMediaAuditResultDetailRequest {
	s.PageNo = &v
	return s
}

func (s *GetMediaAuditResultDetailRequest) Validate() error {
	return dara.Validate(s)
}
