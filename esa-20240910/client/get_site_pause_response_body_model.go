// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSitePauseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPaused(v bool) *GetSitePauseResponseBody
	GetPaused() *bool
	SetRequestId(v string) *GetSitePauseResponseBody
	GetRequestId() *string
}

type GetSitePauseResponseBody struct {
	// Indicates whether ESA is paused on the website. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Paused *bool `json:"Paused,omitempty" xml:"Paused,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSitePauseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSitePauseResponseBody) GoString() string {
	return s.String()
}

func (s *GetSitePauseResponseBody) GetPaused() *bool {
	return s.Paused
}

func (s *GetSitePauseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSitePauseResponseBody) SetPaused(v bool) *GetSitePauseResponseBody {
	s.Paused = &v
	return s
}

func (s *GetSitePauseResponseBody) SetRequestId(v string) *GetSitePauseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSitePauseResponseBody) Validate() error {
	return dara.Validate(s)
}
