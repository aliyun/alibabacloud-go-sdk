// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRowPermissionByTableGuidsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRowPermissionByTableGuidsResponseBody
	GetCode() *string
	SetData(v []*GetRowPermissionByTableGuidsResponseBodyData) *GetRowPermissionByTableGuidsResponseBody
	GetData() []*GetRowPermissionByTableGuidsResponseBodyData
	SetHttpStatusCode(v int32) *GetRowPermissionByTableGuidsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetRowPermissionByTableGuidsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRowPermissionByTableGuidsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRowPermissionByTableGuidsResponseBody
	GetSuccess() *bool
}

type GetRowPermissionByTableGuidsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetRowPermissionByTableGuidsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRowPermissionByTableGuidsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRowPermissionByTableGuidsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRowPermissionByTableGuidsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRowPermissionByTableGuidsResponseBody) GetData() []*GetRowPermissionByTableGuidsResponseBodyData {
	return s.Data
}

func (s *GetRowPermissionByTableGuidsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRowPermissionByTableGuidsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRowPermissionByTableGuidsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRowPermissionByTableGuidsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRowPermissionByTableGuidsResponseBody) SetCode(v string) *GetRowPermissionByTableGuidsResponseBody {
	s.Code = &v
	return s
}

func (s *GetRowPermissionByTableGuidsResponseBody) SetData(v []*GetRowPermissionByTableGuidsResponseBodyData) *GetRowPermissionByTableGuidsResponseBody {
	s.Data = v
	return s
}

func (s *GetRowPermissionByTableGuidsResponseBody) SetHttpStatusCode(v int32) *GetRowPermissionByTableGuidsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRowPermissionByTableGuidsResponseBody) SetMessage(v string) *GetRowPermissionByTableGuidsResponseBody {
	s.Message = &v
	return s
}

func (s *GetRowPermissionByTableGuidsResponseBody) SetRequestId(v string) *GetRowPermissionByTableGuidsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRowPermissionByTableGuidsResponseBody) SetSuccess(v bool) *GetRowPermissionByTableGuidsResponseBody {
	s.Success = &v
	return s
}

func (s *GetRowPermissionByTableGuidsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRowPermissionByTableGuidsResponseBodyData struct {
	// example:
	//
	// Restrict query region
	RowPermissionDesc *string `json:"RowPermissionDesc,omitempty" xml:"RowPermissionDesc,omitempty"`
	// example:
	//
	// 300000001
	RowPermissionId *string `json:"RowPermissionId,omitempty" xml:"RowPermissionId,omitempty"`
	// example:
	//
	// Region
	RowPermissionName *string `json:"RowPermissionName,omitempty" xml:"RowPermissionName,omitempty"`
}

func (s GetRowPermissionByTableGuidsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRowPermissionByTableGuidsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRowPermissionByTableGuidsResponseBodyData) GetRowPermissionDesc() *string {
	return s.RowPermissionDesc
}

func (s *GetRowPermissionByTableGuidsResponseBodyData) GetRowPermissionId() *string {
	return s.RowPermissionId
}

func (s *GetRowPermissionByTableGuidsResponseBodyData) GetRowPermissionName() *string {
	return s.RowPermissionName
}

func (s *GetRowPermissionByTableGuidsResponseBodyData) SetRowPermissionDesc(v string) *GetRowPermissionByTableGuidsResponseBodyData {
	s.RowPermissionDesc = &v
	return s
}

func (s *GetRowPermissionByTableGuidsResponseBodyData) SetRowPermissionId(v string) *GetRowPermissionByTableGuidsResponseBodyData {
	s.RowPermissionId = &v
	return s
}

func (s *GetRowPermissionByTableGuidsResponseBodyData) SetRowPermissionName(v string) *GetRowPermissionByTableGuidsResponseBodyData {
	s.RowPermissionName = &v
	return s
}

func (s *GetRowPermissionByTableGuidsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
