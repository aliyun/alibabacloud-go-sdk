// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlockObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*string) *BlockObjectRequest
	GetContent() []*string
	SetExtension(v string) *BlockObjectRequest
	GetExtension() *string
	SetMaxage(v int32) *BlockObjectRequest
	GetMaxage() *int32
	SetSiteId(v int64) *BlockObjectRequest
	GetSiteId() *int64
	SetType(v string) *BlockObjectRequest
	GetType() *string
}

type BlockObjectRequest struct {
	// The content to block.
	//
	// This parameter is required.
	Content []*string `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The blocking period that you can extend. Set the value to 2year.
	//
	// example:
	//
	// 2year
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The period of time during which the URL is blocked. Unit: seconds. Specify this parameter if Type is set to block.
	//
	// example:
	//
	// 864000
	Maxage *int32 `json:"Maxage,omitempty" xml:"Maxage,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// BlockObject
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The type. Valid values:
	//
	// 	- **block**
	//
	// 	- **unblock**
	//
	// This parameter is required.
	//
	// example:
	//
	// block
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BlockObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s BlockObjectRequest) GoString() string {
	return s.String()
}

func (s *BlockObjectRequest) GetContent() []*string {
	return s.Content
}

func (s *BlockObjectRequest) GetExtension() *string {
	return s.Extension
}

func (s *BlockObjectRequest) GetMaxage() *int32 {
	return s.Maxage
}

func (s *BlockObjectRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *BlockObjectRequest) GetType() *string {
	return s.Type
}

func (s *BlockObjectRequest) SetContent(v []*string) *BlockObjectRequest {
	s.Content = v
	return s
}

func (s *BlockObjectRequest) SetExtension(v string) *BlockObjectRequest {
	s.Extension = &v
	return s
}

func (s *BlockObjectRequest) SetMaxage(v int32) *BlockObjectRequest {
	s.Maxage = &v
	return s
}

func (s *BlockObjectRequest) SetSiteId(v int64) *BlockObjectRequest {
	s.SiteId = &v
	return s
}

func (s *BlockObjectRequest) SetType(v string) *BlockObjectRequest {
	s.Type = &v
	return s
}

func (s *BlockObjectRequest) Validate() error {
	return dara.Validate(s)
}
