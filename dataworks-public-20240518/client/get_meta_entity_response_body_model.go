// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetaEntity(v *MetaEntity) *GetMetaEntityResponseBody
	GetMetaEntity() *MetaEntity
	SetRequestId(v string) *GetMetaEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaEntityResponseBody
	GetSuccess() *bool
}

type GetMetaEntityResponseBody struct {
	MetaEntity *MetaEntity `json:"MetaEntity,omitempty" xml:"MetaEntity,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0A04C673-BEFA-5803-94E5-89E2D9F8C567
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetaEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaEntityResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaEntityResponseBody) GetMetaEntity() *MetaEntity {
	return s.MetaEntity
}

func (s *GetMetaEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaEntityResponseBody) SetMetaEntity(v *MetaEntity) *GetMetaEntityResponseBody {
	s.MetaEntity = v
	return s
}

func (s *GetMetaEntityResponseBody) SetRequestId(v string) *GetMetaEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaEntityResponseBody) SetSuccess(v bool) *GetMetaEntityResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaEntityResponseBody) Validate() error {
	if s.MetaEntity != nil {
		if err := s.MetaEntity.Validate(); err != nil {
			return err
		}
	}
	return nil
}
