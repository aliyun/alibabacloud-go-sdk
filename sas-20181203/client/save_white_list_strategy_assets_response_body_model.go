// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWhiteListStrategyAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveWhiteListStrategyAssetsResponseBody
	GetRequestId() *string
}

type SaveWhiteListStrategyAssetsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveWhiteListStrategyAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveWhiteListStrategyAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWhiteListStrategyAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveWhiteListStrategyAssetsResponseBody) SetRequestId(v string) *SaveWhiteListStrategyAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveWhiteListStrategyAssetsResponseBody) Validate() error {
	return dara.Validate(s)
}
