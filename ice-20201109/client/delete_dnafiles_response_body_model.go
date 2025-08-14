// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDNAFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDNAFilesResponseBody
	GetRequestId() *string
}

type DeleteDNAFilesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDNAFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDNAFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDNAFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDNAFilesResponseBody) SetRequestId(v string) *DeleteDNAFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDNAFilesResponseBody) Validate() error {
	return dara.Validate(s)
}
