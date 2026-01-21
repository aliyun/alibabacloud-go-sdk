// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAutoProtectNewAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v string) *SetAutoProtectNewAssetsResponseBody
	GetModule() *string
	SetRequestId(v string) *SetAutoProtectNewAssetsResponseBody
	GetRequestId() *string
}

type SetAutoProtectNewAssetsResponseBody struct {
	// example:
	//
	// api_server
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// example:
	//
	// F0F82705-CFC7-5F83-86C8-A063892F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAutoProtectNewAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAutoProtectNewAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *SetAutoProtectNewAssetsResponseBody) GetModule() *string {
	return s.Module
}

func (s *SetAutoProtectNewAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAutoProtectNewAssetsResponseBody) SetModule(v string) *SetAutoProtectNewAssetsResponseBody {
	s.Module = &v
	return s
}

func (s *SetAutoProtectNewAssetsResponseBody) SetRequestId(v string) *SetAutoProtectNewAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAutoProtectNewAssetsResponseBody) Validate() error {
	return dara.Validate(s)
}
