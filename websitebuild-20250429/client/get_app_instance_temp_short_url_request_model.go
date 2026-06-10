// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceTempShortUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetAppInstanceTempShortUrlRequest
	GetBizId() *string
}

type GetAppInstanceTempShortUrlRequest struct {
	// Application business ID
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s GetAppInstanceTempShortUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceTempShortUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAppInstanceTempShortUrlRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppInstanceTempShortUrlRequest) SetBizId(v string) *GetAppInstanceTempShortUrlRequest {
	s.BizId = &v
	return s
}

func (s *GetAppInstanceTempShortUrlRequest) Validate() error {
	return dara.Validate(s)
}
