// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRegistryScanDayNumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetRegistryScanDayNumResponseBody
	GetRequestId() *string
}

type SetRegistryScanDayNumResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 48483161-F328-5A12-AB78-3EB81F37****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetRegistryScanDayNumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetRegistryScanDayNumResponseBody) GoString() string {
	return s.String()
}

func (s *SetRegistryScanDayNumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetRegistryScanDayNumResponseBody) SetRequestId(v string) *SetRegistryScanDayNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetRegistryScanDayNumResponseBody) Validate() error {
	return dara.Validate(s)
}
