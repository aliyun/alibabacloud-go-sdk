// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUninstallClientsByUuidsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUninstallClientsByUuidsResponseBody
	GetRequestId() *string
}

type AddUninstallClientsByUuidsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3956048F-9D73-5EDB-834B-4827BB483977
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUninstallClientsByUuidsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUninstallClientsByUuidsResponseBody) GoString() string {
	return s.String()
}

func (s *AddUninstallClientsByUuidsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUninstallClientsByUuidsResponseBody) SetRequestId(v string) *AddUninstallClientsByUuidsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUninstallClientsByUuidsResponseBody) Validate() error {
	return dara.Validate(s)
}
