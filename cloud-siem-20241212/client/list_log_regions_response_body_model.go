// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogRegions(v []*string) *ListLogRegionsResponseBody
	GetLogRegions() []*string
	SetRequestId(v string) *ListLogRegionsResponseBody
	GetRequestId() *string
}

type ListLogRegionsResponseBody struct {
	// A list of log storage region IDs.
	LogRegions []*string `json:"LogRegions,omitempty" xml:"LogRegions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLogRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLogRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogRegionsResponseBody) GetLogRegions() []*string {
	return s.LogRegions
}

func (s *ListLogRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLogRegionsResponseBody) SetLogRegions(v []*string) *ListLogRegionsResponseBody {
	s.LogRegions = v
	return s
}

func (s *ListLogRegionsResponseBody) SetRequestId(v string) *ListLogRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}
