// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeliveryHistoryJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDeliveryHistoryJobRequest
	GetClientToken() *string
	SetTrailName(v string) *CreateDeliveryHistoryJobRequest
	GetTrailName() *string
}

type CreateDeliveryHistoryJobRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests.
	//
	// The token can contain only ASCII characters and can be up to 64 characters in length.
	//
	// For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the trail for which you want to create a historical event delivery task.
	//
	// This parameter is required.
	//
	// example:
	//
	// trail-name
	TrailName *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
}

func (s CreateDeliveryHistoryJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryHistoryJobRequest) GoString() string {
	return s.String()
}

func (s *CreateDeliveryHistoryJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDeliveryHistoryJobRequest) GetTrailName() *string {
	return s.TrailName
}

func (s *CreateDeliveryHistoryJobRequest) SetClientToken(v string) *CreateDeliveryHistoryJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDeliveryHistoryJobRequest) SetTrailName(v string) *CreateDeliveryHistoryJobRequest {
	s.TrailName = &v
	return s
}

func (s *CreateDeliveryHistoryJobRequest) Validate() error {
	return dara.Validate(s)
}
