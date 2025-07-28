// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUrgentDemandItemHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteUrgentDemandItemHeaders
	GetCommonHeaders() map[string]*string
	SetYunUserId(v string) *DeleteUrgentDemandItemHeaders
	GetYunUserId() *string
}

type DeleteUrgentDemandItemHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111222
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s DeleteUrgentDemandItemHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteUrgentDemandItemHeaders) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandItemHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteUrgentDemandItemHeaders) GetYunUserId() *string {
	return s.YunUserId
}

func (s *DeleteUrgentDemandItemHeaders) SetCommonHeaders(v map[string]*string) *DeleteUrgentDemandItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteUrgentDemandItemHeaders) SetYunUserId(v string) *DeleteUrgentDemandItemHeaders {
	s.YunUserId = &v
	return s
}

func (s *DeleteUrgentDemandItemHeaders) Validate() error {
	return dara.Validate(s)
}
