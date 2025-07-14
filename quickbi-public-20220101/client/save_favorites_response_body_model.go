// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveFavoritesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveFavoritesResponseBody
	GetRequestId() *string
	SetResult(v bool) *SaveFavoritesResponseBody
	GetResult() *bool
	SetSuccess(v bool) *SaveFavoritesResponseBody
	GetSuccess() *bool
}

type SaveFavoritesResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the result of the interface execution. Possible values:
	//
	// - true: Execution successful
	//
	// - false: Execution failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: Request successful
	//
	// - false: Request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveFavoritesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveFavoritesResponseBody) GoString() string {
	return s.String()
}

func (s *SaveFavoritesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveFavoritesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *SaveFavoritesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveFavoritesResponseBody) SetRequestId(v string) *SaveFavoritesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveFavoritesResponseBody) SetResult(v bool) *SaveFavoritesResponseBody {
	s.Result = &v
	return s
}

func (s *SaveFavoritesResponseBody) SetSuccess(v bool) *SaveFavoritesResponseBody {
	s.Success = &v
	return s
}

func (s *SaveFavoritesResponseBody) Validate() error {
	return dara.Validate(s)
}
