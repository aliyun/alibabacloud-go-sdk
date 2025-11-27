// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeletedBaksetIds(v *DeleteBackupFileResponseBodyDeletedBaksetIds) *DeleteBackupFileResponseBody
	GetDeletedBaksetIds() *DeleteBackupFileResponseBodyDeletedBaksetIds
	SetRequestId(v string) *DeleteBackupFileResponseBody
	GetRequestId() *string
}

type DeleteBackupFileResponseBody struct {
	// An array that consists of the IDs of deleted backup sets.
	DeletedBaksetIds *DeleteBackupFileResponseBodyDeletedBaksetIds `json:"DeletedBaksetIds,omitempty" xml:"DeletedBaksetIds,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C7B3A91C-0ACD-4948-ACAE-xxxxxxxD4069
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBackupFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupFileResponseBody) GetDeletedBaksetIds() *DeleteBackupFileResponseBodyDeletedBaksetIds {
	return s.DeletedBaksetIds
}

func (s *DeleteBackupFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBackupFileResponseBody) SetDeletedBaksetIds(v *DeleteBackupFileResponseBodyDeletedBaksetIds) *DeleteBackupFileResponseBody {
	s.DeletedBaksetIds = v
	return s
}

func (s *DeleteBackupFileResponseBody) SetRequestId(v string) *DeleteBackupFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupFileResponseBody) Validate() error {
	if s.DeletedBaksetIds != nil {
		if err := s.DeletedBaksetIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteBackupFileResponseBodyDeletedBaksetIds struct {
	DeletedBaksetIds []*int32 `json:"DeletedBaksetIds,omitempty" xml:"DeletedBaksetIds,omitempty" type:"Repeated"`
}

func (s DeleteBackupFileResponseBodyDeletedBaksetIds) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupFileResponseBodyDeletedBaksetIds) GoString() string {
	return s.String()
}

func (s *DeleteBackupFileResponseBodyDeletedBaksetIds) GetDeletedBaksetIds() []*int32 {
	return s.DeletedBaksetIds
}

func (s *DeleteBackupFileResponseBodyDeletedBaksetIds) SetDeletedBaksetIds(v []*int32) *DeleteBackupFileResponseBodyDeletedBaksetIds {
	s.DeletedBaksetIds = v
	return s
}

func (s *DeleteBackupFileResponseBodyDeletedBaksetIds) Validate() error {
	return dara.Validate(s)
}
