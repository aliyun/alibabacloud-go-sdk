// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCdnDomainResponseBody
	GetRequestId() *string
}

type ModifyCdnDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCdnDomainResponseBody) SetRequestId(v string) *ModifyCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
