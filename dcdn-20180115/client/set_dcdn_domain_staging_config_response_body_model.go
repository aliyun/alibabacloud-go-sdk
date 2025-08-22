// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnDomainStagingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDcdnDomainStagingConfigResponseBody
	GetRequestId() *string
}

type SetDcdnDomainStagingConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDcdnDomainStagingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnDomainStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainStagingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDcdnDomainStagingConfigResponseBody) SetRequestId(v string) *SetDcdnDomainStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDcdnDomainStagingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
