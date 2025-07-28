// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveUrgentDemandItemHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SaveUrgentDemandItemHeaders
	GetCommonHeaders() map[string]*string
	SetYunUserId(v string) *SaveUrgentDemandItemHeaders
	GetYunUserId() *string
}

type SaveUrgentDemandItemHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111222
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s SaveUrgentDemandItemHeaders) String() string {
	return dara.Prettify(s)
}

func (s SaveUrgentDemandItemHeaders) GoString() string {
	return s.String()
}

func (s *SaveUrgentDemandItemHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SaveUrgentDemandItemHeaders) GetYunUserId() *string {
	return s.YunUserId
}

func (s *SaveUrgentDemandItemHeaders) SetCommonHeaders(v map[string]*string) *SaveUrgentDemandItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveUrgentDemandItemHeaders) SetYunUserId(v string) *SaveUrgentDemandItemHeaders {
	s.YunUserId = &v
	return s
}

func (s *SaveUrgentDemandItemHeaders) Validate() error {
	return dara.Validate(s)
}
