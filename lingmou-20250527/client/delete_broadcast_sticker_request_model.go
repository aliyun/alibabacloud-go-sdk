// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBroadcastStickerRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteBroadcastStickerRequest struct {
}

func (s DeleteBroadcastStickerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBroadcastStickerRequest) GoString() string {
	return s.String()
}

func (s *DeleteBroadcastStickerRequest) Validate() error {
	return dara.Validate(s)
}
