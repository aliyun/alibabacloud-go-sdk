// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLakebaseS3AccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLakebaseS3AccountResponseBody
	GetRequestId() *string
}

type DeleteLakebaseS3AccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLakebaseS3AccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLakebaseS3AccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLakebaseS3AccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLakebaseS3AccountResponseBody) SetRequestId(v string) *DeleteLakebaseS3AccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLakebaseS3AccountResponseBody) Validate() error {
	return dara.Validate(s)
}
