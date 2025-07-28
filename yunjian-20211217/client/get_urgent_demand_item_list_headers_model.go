// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUrgentDemandItemListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUrgentDemandItemListHeaders
	GetCommonHeaders() map[string]*string
	SetYunUserId(v string) *GetUrgentDemandItemListHeaders
	GetYunUserId() *string
}

type GetUrgentDemandItemListHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111222
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s GetUrgentDemandItemListHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandItemListHeaders) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUrgentDemandItemListHeaders) GetYunUserId() *string {
	return s.YunUserId
}

func (s *GetUrgentDemandItemListHeaders) SetCommonHeaders(v map[string]*string) *GetUrgentDemandItemListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUrgentDemandItemListHeaders) SetYunUserId(v string) *GetUrgentDemandItemListHeaders {
	s.YunUserId = &v
	return s
}

func (s *GetUrgentDemandItemListHeaders) Validate() error {
	return dara.Validate(s)
}
