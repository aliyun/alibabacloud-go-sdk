// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizTraceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItem(v *BizTraceConfig) *GetBizTraceResponseBody
	GetItem() *BizTraceConfig
	SetRequestId(v string) *GetBizTraceResponseBody
	GetRequestId() *string
}

type GetBizTraceResponseBody struct {
	Item *BizTraceConfig `json:"item,omitempty" xml:"item,omitempty"`
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetBizTraceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBizTraceResponseBody) GoString() string {
	return s.String()
}

func (s *GetBizTraceResponseBody) GetItem() *BizTraceConfig {
	return s.Item
}

func (s *GetBizTraceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBizTraceResponseBody) SetItem(v *BizTraceConfig) *GetBizTraceResponseBody {
	s.Item = v
	return s
}

func (s *GetBizTraceResponseBody) SetRequestId(v string) *GetBizTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBizTraceResponseBody) Validate() error {
	if s.Item != nil {
		if err := s.Item.Validate(); err != nil {
			return err
		}
	}
	return nil
}
