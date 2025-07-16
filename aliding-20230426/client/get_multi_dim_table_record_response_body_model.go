// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedBy(v *GetMultiDimTableRecordResponseBodyCreatedBy) *GetMultiDimTableRecordResponseBody
	GetCreatedBy() *GetMultiDimTableRecordResponseBodyCreatedBy
	SetCreatedTime(v int64) *GetMultiDimTableRecordResponseBody
	GetCreatedTime() *int64
	SetFields(v map[string]interface{}) *GetMultiDimTableRecordResponseBody
	GetFields() map[string]interface{}
	SetId(v string) *GetMultiDimTableRecordResponseBody
	GetId() *string
	SetLastModifiedBy(v *GetMultiDimTableRecordResponseBodyLastModifiedBy) *GetMultiDimTableRecordResponseBody
	GetLastModifiedBy() *GetMultiDimTableRecordResponseBodyLastModifiedBy
	SetLastModifiedTime(v int64) *GetMultiDimTableRecordResponseBody
	GetLastModifiedTime() *int64
	SetRequestId(v string) *GetMultiDimTableRecordResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetMultiDimTableRecordResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetMultiDimTableRecordResponseBody
	GetVendorType() *string
}

type GetMultiDimTableRecordResponseBody struct {
	CreatedBy *GetMultiDimTableRecordResponseBodyCreatedBy `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// xxx
	Fields map[string]interface{} `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// example:
	//
	// xxx
	Id             *string                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	LastModifiedBy *GetMultiDimTableRecordResponseBodyLastModifiedBy `json:"LastModifiedBy,omitempty" xml:"LastModifiedBy,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	LastModifiedTime *int64 `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetMultiDimTableRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableRecordResponseBody) GetCreatedBy() *GetMultiDimTableRecordResponseBodyCreatedBy {
	return s.CreatedBy
}

func (s *GetMultiDimTableRecordResponseBody) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetMultiDimTableRecordResponseBody) GetFields() map[string]interface{} {
	return s.Fields
}

func (s *GetMultiDimTableRecordResponseBody) GetId() *string {
	return s.Id
}

func (s *GetMultiDimTableRecordResponseBody) GetLastModifiedBy() *GetMultiDimTableRecordResponseBodyLastModifiedBy {
	return s.LastModifiedBy
}

func (s *GetMultiDimTableRecordResponseBody) GetLastModifiedTime() *int64 {
	return s.LastModifiedTime
}

func (s *GetMultiDimTableRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultiDimTableRecordResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetMultiDimTableRecordResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetMultiDimTableRecordResponseBody) SetCreatedBy(v *GetMultiDimTableRecordResponseBodyCreatedBy) *GetMultiDimTableRecordResponseBody {
	s.CreatedBy = v
	return s
}

func (s *GetMultiDimTableRecordResponseBody) SetCreatedTime(v int64) *GetMultiDimTableRecordResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetMultiDimTableRecordResponseBody) SetFields(v map[string]interface{}) *GetMultiDimTableRecordResponseBody {
	s.Fields = v
	return s
}

func (s *GetMultiDimTableRecordResponseBody) SetId(v string) *GetMultiDimTableRecordResponseBody {
	s.Id = &v
	return s
}

func (s *GetMultiDimTableRecordResponseBody) SetLastModifiedBy(v *GetMultiDimTableRecordResponseBodyLastModifiedBy) *GetMultiDimTableRecordResponseBody {
	s.LastModifiedBy = v
	return s
}

func (s *GetMultiDimTableRecordResponseBody) SetLastModifiedTime(v int64) *GetMultiDimTableRecordResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *GetMultiDimTableRecordResponseBody) SetRequestId(v string) *GetMultiDimTableRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiDimTableRecordResponseBody) SetVendorRequestId(v string) *GetMultiDimTableRecordResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetMultiDimTableRecordResponseBody) SetVendorType(v string) *GetMultiDimTableRecordResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetMultiDimTableRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMultiDimTableRecordResponseBodyCreatedBy struct {
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetMultiDimTableRecordResponseBodyCreatedBy) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableRecordResponseBodyCreatedBy) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableRecordResponseBodyCreatedBy) GetUserId() *string {
	return s.UserId
}

func (s *GetMultiDimTableRecordResponseBodyCreatedBy) SetUserId(v string) *GetMultiDimTableRecordResponseBodyCreatedBy {
	s.UserId = &v
	return s
}

func (s *GetMultiDimTableRecordResponseBodyCreatedBy) Validate() error {
	return dara.Validate(s)
}

type GetMultiDimTableRecordResponseBodyLastModifiedBy struct {
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetMultiDimTableRecordResponseBodyLastModifiedBy) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableRecordResponseBodyLastModifiedBy) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableRecordResponseBodyLastModifiedBy) GetUserId() *string {
	return s.UserId
}

func (s *GetMultiDimTableRecordResponseBodyLastModifiedBy) SetUserId(v string) *GetMultiDimTableRecordResponseBodyLastModifiedBy {
	s.UserId = &v
	return s
}

func (s *GetMultiDimTableRecordResponseBodyLastModifiedBy) Validate() error {
	return dara.Validate(s)
}
