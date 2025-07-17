// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataArchiveCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderResultType(v string) *GetDataArchiveCountRequest
	GetOrderResultType() *string
	SetPluginType(v string) *GetDataArchiveCountRequest
	GetPluginType() *string
	SetSearchDateType(v string) *GetDataArchiveCountRequest
	GetSearchDateType() *string
	SetTid(v int64) *GetDataArchiveCountRequest
	GetTid() *int64
}

type GetDataArchiveCountRequest struct {
	// The type of the identity. Default value: AS_ADMIN.
	//
	// example:
	//
	// AS_ADMIN
	OrderResultType *string `json:"OrderResultType,omitempty" xml:"OrderResultType,omitempty"`
	// The plugin type. Default value: DATA_ARCHIVE.
	//
	// example:
	//
	// DATA_ARCHIVE
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	// The time when the ticket is modified or created. The statistics of data archiving tickets are calculated based on the creation time.
	//
	// example:
	//
	// CREATE_TIME
	SearchDateType *string `json:"SearchDateType,omitempty" xml:"SearchDateType,omitempty"`
	// The tenant ID.
	//
	// >  To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 2****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataArchiveCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveCountRequest) GoString() string {
	return s.String()
}

func (s *GetDataArchiveCountRequest) GetOrderResultType() *string {
	return s.OrderResultType
}

func (s *GetDataArchiveCountRequest) GetPluginType() *string {
	return s.PluginType
}

func (s *GetDataArchiveCountRequest) GetSearchDateType() *string {
	return s.SearchDateType
}

func (s *GetDataArchiveCountRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataArchiveCountRequest) SetOrderResultType(v string) *GetDataArchiveCountRequest {
	s.OrderResultType = &v
	return s
}

func (s *GetDataArchiveCountRequest) SetPluginType(v string) *GetDataArchiveCountRequest {
	s.PluginType = &v
	return s
}

func (s *GetDataArchiveCountRequest) SetSearchDateType(v string) *GetDataArchiveCountRequest {
	s.SearchDateType = &v
	return s
}

func (s *GetDataArchiveCountRequest) SetTid(v int64) *GetDataArchiveCountRequest {
	s.Tid = &v
	return s
}

func (s *GetDataArchiveCountRequest) Validate() error {
	return dara.Validate(s)
}
