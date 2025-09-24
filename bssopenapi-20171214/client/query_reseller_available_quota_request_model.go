// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResellerAvailableQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItemCodes(v string) *QueryResellerAvailableQuotaRequest
	GetItemCodes() *string
	SetOwnerId(v int64) *QueryResellerAvailableQuotaRequest
	GetOwnerId() *int64
}

type QueryResellerAvailableQuotaRequest struct {
	ItemCodes *string `json:"ItemCodes,omitempty" xml:"ItemCodes,omitempty"`
	// This parameter is required.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s QueryResellerAvailableQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryResellerAvailableQuotaRequest) GoString() string {
	return s.String()
}

func (s *QueryResellerAvailableQuotaRequest) GetItemCodes() *string {
	return s.ItemCodes
}

func (s *QueryResellerAvailableQuotaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryResellerAvailableQuotaRequest) SetItemCodes(v string) *QueryResellerAvailableQuotaRequest {
	s.ItemCodes = &v
	return s
}

func (s *QueryResellerAvailableQuotaRequest) SetOwnerId(v int64) *QueryResellerAvailableQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryResellerAvailableQuotaRequest) Validate() error {
	return dara.Validate(s)
}
