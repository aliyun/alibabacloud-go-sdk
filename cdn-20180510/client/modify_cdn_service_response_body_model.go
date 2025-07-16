// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdnServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCdnServiceResponseBody
	GetRequestId() *string
}

type ModifyCdnServiceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCdnServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdnServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCdnServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCdnServiceResponseBody) SetRequestId(v string) *ModifyCdnServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCdnServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
