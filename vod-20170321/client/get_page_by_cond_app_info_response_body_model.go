// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageByCondAppInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetPageByCondAppInfoResponseBodyData) *GetPageByCondAppInfoResponseBody
	GetData() *GetPageByCondAppInfoResponseBodyData
	SetRequestId(v string) *GetPageByCondAppInfoResponseBody
	GetRequestId() *string
}

type GetPageByCondAppInfoResponseBody struct {
	Data      *GetPageByCondAppInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPageByCondAppInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPageByCondAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPageByCondAppInfoResponseBody) GetData() *GetPageByCondAppInfoResponseBodyData {
	return s.Data
}

func (s *GetPageByCondAppInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPageByCondAppInfoResponseBody) SetData(v *GetPageByCondAppInfoResponseBodyData) *GetPageByCondAppInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetPageByCondAppInfoResponseBody) SetRequestId(v string) *GetPageByCondAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPageByCondAppInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPageByCondAppInfoResponseBodyData struct {
	List  []*AppInfoDTO `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Total *int64        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetPageByCondAppInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPageByCondAppInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPageByCondAppInfoResponseBodyData) GetList() []*AppInfoDTO {
	return s.List
}

func (s *GetPageByCondAppInfoResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetPageByCondAppInfoResponseBodyData) SetList(v []*AppInfoDTO) *GetPageByCondAppInfoResponseBodyData {
	s.List = v
	return s
}

func (s *GetPageByCondAppInfoResponseBodyData) SetTotal(v int64) *GetPageByCondAppInfoResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetPageByCondAppInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
