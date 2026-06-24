// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineAppInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *OfflineAppInstanceRequest
	GetBizId() *string
}

type OfflineAppInstanceRequest struct {
	// The business ID of the application instance.
	//
	// example:
	//
	// WD20250814102215000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s OfflineAppInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *OfflineAppInstanceRequest) GetBizId() *string {
	return s.BizId
}

func (s *OfflineAppInstanceRequest) SetBizId(v string) *OfflineAppInstanceRequest {
	s.BizId = &v
	return s
}

func (s *OfflineAppInstanceRequest) Validate() error {
	return dara.Validate(s)
}
